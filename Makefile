# Setting it to current path if env doesn't exist
# This is used when running "make test" in local
ifeq ($(WORKSPACE),)
        WORKSPACE=.
endif

test:
	@echo ">>-> Running go test for the libraries"; \
        echo "Running UT for library: '$(lib)'"; \
        go1.23.5 test -v ./... -coverprofile=coverage.txt -covermode count --cover -coverpkg=./... -p 1 && \
		gocover-cobertura -ignore-dirs '/mocks$$' -ignore-files '(main.go$$|config.go$$|^internal/testdependencies/manager.go$$)' < coverage.txt > ${WORKSPACE}/$(lib)_coverage.xml || { \
		echo "test failed for $${entry}."; \
		exit 1; \
        }; \

# DO NOT CHANGE, misplacing `-delete` argument might delete
# everything in the folder.
clean:
	 find . -name "*coverage.xml" -type f -delete
	 find . -name "*coverage.txt" -type f -delete
