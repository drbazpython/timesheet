hello:
	echo "Hello"

buildTS:
	@echo "Building .... "
	@go build -o ./bin drbaz.com/timesheet/cmd/timesheet
	@echo "Installing ...."
	@go install drbaz.com/timesheet/cmd/timesheet
	@echo "Timesheet built to bin and installed"
	@echo "Running timesheet ...."
	@timesheet
	
run:
	go run cmd/timesheet/main.go

tidy:
	go mod tidy
	
	