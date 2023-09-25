# PrePush
A CLI tool to run scripts or commands before your git push

PrePush is a small but powerful tool designed to make your development process smoother. This tool helps developers to automate the execution of any pre-defined commands before a commit is made ready to push to the git repository.

## Features

- Automates the running of pre-defined commands before a commit.
- Easy to use.
- Helps in maintaining high code quality by potentially preventing problematic commits.
- Both JSON and YAML configuration files are supported.

## Installation
```
Coming Soon...
```

### Configuration
PrePush can be configured using a configuration file. The configuration file can be either in JSON or YAML format.
The configuration file should be named as `prepush.json` or `prepush.yaml` or `prepush.yml` and should be placed in the root directory of the project.

#### JSON Configuration

PrePush can make the configuration file for you. Just run the following command in the root directory of your project.
```bash
prepush init
```

#### JSON Configuration
```json
{
    "config" : {
      "name" : "value"
    },

    "commands": {
      "Hello PrePush!" :  "echo \"Hello PrePush\"",
      "Listing current directory" :  "ls -l"
    }
}
```

#### YAML Configuration
```yaml
config:
  name: value

commands:
  Hello PrePush!: echo "Hello PrePush"
  Listing current directory: ls -l
```

## Usage
#### run your commands
```
prepush
```

## Contributing

```
Coming Soon...
```

## License

This project is licensed under the MPL-2.0 license - see the `LICENSE.md` file for details.