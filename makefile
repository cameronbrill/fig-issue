frontend-start:
	cd frontend && doppler run -- yarn s

frontend-build:
	cd frontend && doppler run -- yarn build

backend-start:
	cd backend && doppler run -- air