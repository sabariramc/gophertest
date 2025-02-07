# Setting it to current path if env doesn't exist
# This is used when running "make test" in local
ifeq ($(WORKSPACE),)
        WORKSPACE=.
endif

test:
	@echo ">>-> Running go test for the libraries"; \
        echo "Running UT for library: '$(lib)'"; \
        go1.23.5 test -v ./... -coverprofile=coverage.txt -covermode count --cover -coverpkg=./... -p 1 && \
		gocover-cobertura -ignore-dirs '/mocks$$' -ignore-files '(main.go$$|^cmd/policy-downloader/bootstrap/metrics.go$$|^cmd/policy-downloader/bootstrap/server.go$$|config.go$$|^comptest/test_dependency_manager.go$$|store_redis_client.go$$|stream_redis_client.go$$|redis_client.go$$|vault_client.go$$)' < coverage.txt > ${WORKSPACE}/$(lib)_coverage.xml || { \
		echo "test failed for $${entry}."; \
		exit 1; \
        }; \

# DO NOT CHANGE, misplacing `-delete` argument might delete
# everything in the folder.
clean:
	 find . -name "*coverage.xml" -type f -delete
	 find . -name "*coverage.txt" -type f -delete
