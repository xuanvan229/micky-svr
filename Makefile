
build:
	docker build -t micky-svr .

up:
	docker-compose up --abort-on-container-exit --remove-orphans mickyapp mongo_micky

backup-db:
	docker-compose -f docker-compose.yml run --rm -T backup-db | pv > backup.tar.bz2
	# docker-compose -f docker-compose.yml run --rm -T backup-db

restore_db:
	cat backup.tar.bz2 | pv | docker-compose -f docker-compose.yml run --rm -T restore-db