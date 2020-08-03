install:
	@npm install -g coffee-script@1.6.2 node-sass
	@go get github.com/lenghan4real/train
	@go get github.com/revel/cmd/revel
	@go build -o ./bin/train github.com/lenghan4real/train/cmd
server:
	revel run 
release: assets
	revel -v package -a github.com/lenghan4real/mediom
assets:
	@train --source app/assets --out public
test:
	@cd app; go test
	@cd app/controllers; go test
	@cd app/models; go test
	@cd tests; go test
