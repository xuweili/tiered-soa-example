base = $(PWD)

all:
	cd $(base)/zeus && go build
	cd $(base)/athena && go build
	cd $(base)/aphrodite && go build
	cd $(base)/dionysus && go build
	cd $(base)/hercules && go build

run:
	cd $(base)/zeus && ./zeus &
	cd $(base)/athena && ./athena &
	cd $(base)/aphrodite && ./aphrodite &
	cd $(base)/dionysus && ./dionysus &
	cd $(base)/hercules && ./hercules &

kill:
	killall zeus
	killall athena
	killall aphrodite
	killall dionysus
	killall hercules

clean:
	rm -rf zeus/zeus
	rm -rf athena/athena
	rm -rf aphrodite/aphrodite
	rm -rf dionysus/dionysus
	rm -rf hercules/hercules
