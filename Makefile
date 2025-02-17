bin         := target/bin
go_src := $(shell find * -name *.go -not -path "$(vendor)/*" -not -path "$(target)/*")
go_out := $(patsubst cmd/%/main.go,$(bin)/%,$(wildcard cmd/*/main.go))


build: $(go_out)

$(bin)/%: cmd/%/main.go $(go_src) | $(bin)
	@go build --trimpath -o=$@ $<

$(bin):
	@mkdir -p $@

check-latency:
	bash -c 'set -a; . ./.env; set +a; go run cmd/latency/main.go run'