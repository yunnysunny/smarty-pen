branch := $(GIT_BRANCH_FOR_MAKE)
now := $(shell date '+%Y%m%d%H%M%S')
EMPTY :=
projectName := drawer

ifeq ($(branch),$(EMPTY))
	branch := test
endif


all: pull run

check:	

# 需要首先安装依赖包  sudo apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev
# yum install SDL2{,_image,_mixer,_ttf,_gfx}-devel
# brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
# pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_image,_mixer,_ttf,_gfx}
build:
	go build -o ./bin/$(projectName)

# 需要安装依赖 sudo apt install graphviz -y
# go get github.com/kisielk/godepgraph
dep:
	cd src; godepgraph -s github.com/smarty-pen/$(projectName) | dot -Tpng -o ../coverage/godepgraph.png


pull:check
	git checkout $(branch) && git pull origin $(branch)

test:check
	mkdir -p coverage && cd src/ && \
	go test ./... -v -timeout 20s -convey-story -cover -coverprofile=../coverage/coverage.out

coverage:test
	cd src && \
	go tool cover -func ../coverage/coverage.out && \
	go tool cover -html ../coverage/coverage.out -o ../coverage/index.html && \
	scp -pr ../coverage $(COVERAGE_REPORT_DEST)

run:check install
	


clean:
	rm -rf bin/*

grace:pull test run
	
.PHONY: check  pull test coverage run clean  build dep grace
