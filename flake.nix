{
  description = "go devshell and package, created by scaffolder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "go-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
            just
          ];
        };

        packages.gh-nixpkgs-review = pkgs.buildGoModule {
          pname = "gh-nixpkgs-review";
          version = "0.1.1";

          src = self;

          vendorHash = "sha256-X9lokE0GThwhMxmAjeKa7dkMvFhhN7/w/1bjKUr+jdI=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A very small gh extension to provide a more convenient way to use nixpkgs-review-gha";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.gh-nixpkgs-review = {
          type = "app";
          program = "${self.packages.${pkgs.stdenv.hostPlatform.system}.gh-nixpkgs-review}/bin/gh-nixpkgs-review";
        };
      });
}
