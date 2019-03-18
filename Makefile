.PHONY: default
default:
	@echo there is no default

.PHONY: test
test:
	go test
	cd cmd/imgadm_proxy && go test
	cd cmd/vmadm_proxy && go test
