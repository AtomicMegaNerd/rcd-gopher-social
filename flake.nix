{
  description =
    "This is a sample app that is part of the Backend Engineering with Go course on Udemy";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in {
        devShell = pkgs.mkShell {
          # The packages we need for this project
          buildInputs = with pkgs; [
            go_1_23
            go-tools
            gopls
            golangci-lint
            go-task
            air
            rainfrog
            docker-compose
          ];
        };
      });
}
