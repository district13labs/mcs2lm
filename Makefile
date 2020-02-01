.PHONY: setup-githooks

setup-githooks:
	@ln -s ${PWD}/scripts/githooks/precommit/pre-commit ${PWD}/.git/hooks/pre-commit
