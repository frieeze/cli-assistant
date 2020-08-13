# CLI Assistant

CLI Assistant is a command line application that allow to create and execute commands

## Add command
`add` command is used to register a new command. Notice it will overrite the command if the `Command Name` already exist.

There is 2 types of commands : `application` and `browser`.

`application` run console command ( ex: `ls` ).

`browser` open a new tab in your default browser with the link provided.

The assistant will ask you a few informations to build a new command :

-   `Command Name`  : what you have to enter to execute this new command
-   `Command Type`  : Choose between `application` and `browser`. Default is `application`
-   `Command Alias` : If `application` enter the alias of the bash command (ex: for `ls -la /home` enter `ls`). If `browser` enter the URL.
-   `Command Args`  : All your command arguments (ex: for `ls -la /home` enter `-la /home`). Useless if `browser`
-   `Command Description`   : Description of the new command

## Remove command
`del` removes a registered command.