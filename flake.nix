{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    pyproject-nix = {
      url = "github:nix-community/pyproject.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    nixpkgs-terraform.url = "github:stackbuilders/nixpkgs-terraform";
  };

  nixConfig = {
    extra-substituters = "https://nixpkgs-terraform.cachix.org";
    extra-trusted-public-keys = "nixpkgs-terraform.cachix.org-1:8Sit092rIdAVENA3ZVeH9hzSiqI/jng6JiCrQ1Dmusw=";
  };

  outputs = { nixpkgs, flake-utils, pyproject-nix, nixpkgs-terraform, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        terraform = nixpkgs-terraform.packages.${system}."1.5.7";
      in
      {
        devShells.default = pkgs.mkShell {
          packages = [
            (
              pkgs.python3.pkgs.buildPythonApplication
                (
                  (pyproject-nix.lib.renderers.buildPythonPackage {
                    project = (pyproject-nix.lib.project.loadPoetryPyproject {
                      projectRoot = (pkgs.fetchFromGitHub {
                        owner = "cruft";
                        repo = "cruft";
                        rev = "2.15.0";
                        hash = "sha256-qIVyNMoI3LsoOV/6XPa60Y1vTRvkezesF7wF9WVSLGk=";
                      });
                    });
                    python = pkgs.python3;
                  })
                )

            )
            pkgs.go
            pkgs.unzip
            terraform
          ];
        };
      });
}
