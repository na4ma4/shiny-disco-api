-include .makefiles/Makefile
-include .makefiles/pkg/protobuf/v2/Makefile
-include .makefiles/pkg/go/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

######################
# Custom
######################

artifacts/protobuf/go.proto_paths.jq: artifacts/protobuf/args/go
	-@mkdir -p "$(MF_PROJECT_ROOT)/$(@D)"
	jq -Rn 'inputs | select(.)' < "$(^)" > "$(@)"

.vscode/settings.json: artifacts/protobuf/go.proto_paths.jq
	-@mkdir -p "$(MF_PROJECT_ROOT)/$(@D)"
	$(if $(shell cat "$(@)" 2>/dev/null),cat "$(@)",echo '{}') | jq --slurpfile po "$(<)" '.protoc.options=$$po' > "$(@).tmp"
	mv "$(@).tmp" "$(@)"

.PHONY: update-paths
update-paths: .vscode/settings.json
