run:
	npx tailwind \
		-i './assets/css/input.css' \
		-o './assets/css/output.css' \
		--watch & \
	air -c ./.air.toml

build:
	echo "build"

up: 
	echo "up"

down:
	echo "down"

seed: 
	echo "seed"

drop:
	echo "drop"

reset:
	echo "reset"

