{ lib
, buildGoModule
, nixosTests
, testers
, installShellFiles
}:
let
  version = "1.0.1";
  owner = "bezahl-online";
  repo = "ptapi";
  rev = "v${version}";
  sha256 = "1ikyp1rxrl8lyfbll501f13yir1axighnr8x3ji3qzwin6i3w497";
in
buildGoModule {
  pname = "ptapiserver";
  inherit version;

  src = ./.;
 
  vendorSha256 = "";

  buildPhase = ''
    runHook preBuild
    CGO_ENABLED=0 go build -o ptapiserver .
    runHook postBuild
  '';

  installPhase = ''
    mkdir -p $out/bin
    mv ptapiserver $out/bin
    cp localhost.crt localhost.key $out/bin
  '';

  meta = with lib; {
    homepage = "https://github.com/bezahl-online/ptapi";
    description = "ptapi server code";
    license = licenses.mit;
    maintainers = with maintainers; [ /* list of maintainers here */ ];
  };
}

