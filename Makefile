.PHONY: build
build: clean prepare
	@echo "you can use output/bootstrap.sh to boost your application"

.PHONY: clean
clean:
	@rm -rf output/
	@echo "clean done"

.PHONY: prepare
prepare:
	@echo "start to prepare"
	@mkdir -p output/bin output/conf
	@find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/
	@protoc --go_out=. domain/*.proto
	@echo "prepare done"
