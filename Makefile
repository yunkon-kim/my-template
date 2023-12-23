default:
	cd cmd/my-template && $(MAKE)

cc:
	cd cmd/my-template && $(MAKE)

run:
	cd cmd/my-template && $(MAKE) run

runwithport:
	cd cmd/my-template && $(MAKE) runwithport --port=$(PORT)

clean:
	cd cmd/my-template && $(MAKE) clean

prod:
	cd cmd/my-template && $(MAKE) prod

swag swagger:
	cd pkg/ && $(MAKE) swag
