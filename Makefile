.PHONY: build
build: clean
	@echo "building complete"

.PHONY: clean
clean:
	@rm -rf output/
	@echo "clean done"