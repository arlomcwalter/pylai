# Pylai
A simple cli totp manager and authenticator.  

## Usage
Use `pylai --help` to see the available commands.  
Use `pylai [command] --help` for more information about a command.  
Use `pylai --version` to see the current version.  

### Commands:
```
add         Add a new account.
delete      Delete an account.
export      Exports an existing account's secret.
export-all  Exports the entire database.
generate    Generate a TOTP code.
list        List all accounts.
verify      Verifies that a TOTP code matches the account.
```

## Installation
Binaries are available on the [releases page](https://github.com/arlomcwalter/pylai/releases/latest).    
You can also install pylai using homebrew:  
```
brew tap arlomcwalter/tap
brew install pylai
```
