DEPDIR = .deps
MAKEDEPEND = echo "$@: $$(go list -f '{{ join .Deps "\n" }}' $< | \
	awk '/github.com\/prantlf\/go-upload-demo/ { gsub(/^github.com\/[a-z]*\/[-a-z]*\//, ""); printf $$0"/*.go " }')" > $(DEPDIR)/$@.d

TARGETS = serve-echo upload-buffered upload-piped-chunked upload-piped-sized \
	upload-composed-chunked upload-composed-sized

all: vet binaries

binaries: $(TARGETS)

vet:
	go vet ./...

clean:
	go clean
	rm -rf $(TARGETS) $(DEPDIR)

$(DEPDIR)/%.d: ;
.PRECIOUS: $(DEPDIR)/%.d

$(DEPDIR):
	mkdir -p $@

include $(patsubst %,$(DEPDIR)/%.d,$(TARGETS))

%: cmd/%/main.go $(DEPDIR) $(DEPDIR)/%.d
	$(MAKEDEPEND)
	go build -o $@ $<

.PHONY: binaries vet test clean

.DEFAULT_GOAL = binaries
