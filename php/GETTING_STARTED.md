# Getting Started with Slow Code kata in PHP

## Prerequisites

- macOS, Linux or Windows
- [git](https://git-scm.com/) client
- [curl](https://curl.se/download.html) command line utility
- [PHP](https://www.php.net/manual/en/install.php)
    <details><summary>PHP version</summary>

    The kata is configured to run with PHP version 8.1 or above.
  
    To use a different version, simply update the value for `php` in [composer.json](composer.json).
    </details>
- [Composer](https://getcomposer.org/)

## Getting ready

### 1 - Clone the kata repository
```shell
git clone https://github.com/philou/Kata-Slow-Code.git
```
### 2 - Go to the kata's `php` directory
```shell
cd Kata-Slow-Code/php
```
### 3 setup environment with Composer
```shell
composer update
```

## Running the kata

You can run the kata from the command line or from your IDE of choice.

You may also run it using [TCR](../tcr/TCR.md) if you want to add a bit of spice.

- [From a terminal with PHUUnit](#running-the-kata-from-a-terminal)
- [From a terminal with TCR](#running-the-kata-from-a-terminal-with-tcr)
- [From PyCharm](#running-the-kata-from-pycharm)
- [From PyCharm with TCR](#running-the-kata-from-pycharm-with-tcr)
- [From IntelliJ](#running-the-kata-from-intellij)
- [From IntelliJ with TCR](#running-the-kata-from-intellij-with-tcr)
- [From Visual Studio Code](#running-the-kata-from-visual-studio-code)
- [From Visual Studio Code with TCR](#running-the-kata-from-visual-studio-code-with-tcr)

<a name="running-the-kata-from-a-terminal"/></a>

### Running the kata from a terminal

> ***Reminder***: the commands below should be run from [Kata-Slow-Code/php](.) directory

To run the tests:

```shell
phpunit tests
 ```

<a name="running-the-kata-from-a-terminal-with-tcr"/></a>

### Running the kata from a terminal with TCR

> ***Note to Windows users***
>
> Use a **git bash** terminal for running the command below.
> _Windows CMD and PowerShell are not supported_

> ***Reminder***: the commands below should be run from [Kata-Slow-Code/python](.) directory

To create and start the python virtual environment:

Type the following to start TCR:

```shell
./tcrw
```

Refer to [Using TCR](#using-tcr) section for additional details about TCR and available options.

<a name="running-the-kata-from-webstorm"/></a>

### Running the kata from WebStorm

Open WebStorm and select:

`File` > `Open` > `Kata-Slow-Code` > `php`

Run all the tests in the project.

The "**Run**" tool window should display all the executed tests.


<a name="running-the-kata-from-webstorm-with-tcr"/></a>

### Running the kata from WebStorm with TCR

TCR is provided as a command line utility running in a terminal.
You can run it from WebStorm directly, through leveraging on its built-in terminal.

#### 1. Open the kata

Open WebStorm and select:

`File` > `Open` > `Kata-Slow-Code` > `php`

#### 2. Turn off auto-save

TCR is constantly watching the filesystem for changes.
For this reason you need to turn off WebStorm's auto-save in order for it to behave as expected.

`File` > `Settings` > `Appearance & Behavior` > `System Settings`

Under `Autosave` section, uncheck the 2 following options:

- [ ] Save files if the IDE is idle for ___ seconds
- [ ] Save files when switching to a different application or a built-in terminal

#### 3. Configure the built-in terminal to run git bash

> ***Windows Only***
>
> Skip this step if you're on macOS or Linux

PyCharm for Windows is usually set up to run PowerShell by default in its built-in terminal.
TCR does not run in PowerShell.

`File` > `Settings` > `Tools` > `Terminal`

Under `Application Settings` section, set the `Shell path` to `C:\Program Files\Git\bin\bash.exe`

The above path is for a default git installation location. You may need to adjust it in case you have installed git at a
different location.

#### 4. Open a built-in terminal

`View` > `Tool Windows` > `Terminal`

#### 5. Launch TCR

> ***Reminder***: the command below should be run from [Kata-Slow-Code/php](.) directory

From the built-in terminal:

```shell
./tcrw
```

Refer to [Using TCR](#using-tcr) section for additional details about TCR and available options.

<a name="running-the-kata-from-intellij"/></a>

### Running the kata from IntelliJ

> ***Important***: This requires to have `IntelliJ IDEA Ultimate` edition.
>
> IntelliJ's PHP plugin is not available with the Community edition.
> You can still use it as an editor for php files, but you will not
> be able to use the benefits brought by the php plugin, such as syntax highlighting,
> code refactoring or integration with IntelliJ's test navigator.

#### 1. Install `PHP plugin`

- Open IntelliJ and select:`File` > `Settings` > `Plugins` > `Marketplace`
- Search for `PHP` and click `Install`
- Restart `IntelliJ` to activate the plugin

#### 2. Launch IntelliJ and open the project

`File` > `Open` > `Kata-Slow-Code` > `php`

#### 3. Configure the project's PHP SDK [untested]

`File` > `Project Structure` > `Project` > `SDK` > `PHP 8.x` > `OK`

***Note*** If no `PHP SDK` is proposed, select `Add SDK` > `PHP SDK...`
from the SDK dropdown link to tell IntelliJ where it can find PHP on your machine.

#### 4. Configure the project source root directory

From the `Project` navigator window, right click on `src` directory and select
`Mark directory as` > `Sources Root`.

#### 5. Run the tests
 
From the `Project` navigator window, right click on `tests` directory and select
`Run PHP tests in test...`.

The "**Run**" tool window should display all the executed tests.

<a name="running-the-kata-from-intellij-with-tcr"/></a>

### Running the kata from IntelliJ with TCR

> ***Important***: This requires to have `IntelliJ IDEA Ultimate` edition.
>
> Refer to the section [Running the kata from IntelliJ](#running-the-kata-from-intellij)
> for instructions related to installation of IntelliJ's PHP plugin.

TCR is provided as a command line utility running in a terminal.
You can run it from IntelliJ directly, through leveraging on its built-in terminal.

#### 1. Open the kata

Open IntelliJ and select:

`File` > `Open` > `Kata-Slow-Code` > `php`

#### 2. Turn off auto-save

TCR is constantly watching the filesystem for changes.
For this reason you need to turn off IntelliJ's auto-save in order for it to behave as expected.

`File` > `Settings` > `Appearance & Behavior` > `System Settings`

Under `Autosave` section, uncheck the 2 following options:

- [ ] Save files if the IDE is idle for ___ seconds
- [ ] Save files when switching to a different application or a built-in terminal

#### 3. Configure the built-in terminal to run git bash

> ***Windows Only***
>
> Skip this step if you're on macOS or Linux

IntelliJ for Windows is usually set up to run PowerShell by default in its built-in terminal.
TCR does not run in PowerShell.

`File` > `Settings` > `Tools` > `Terminal`

Under `Application Settings` section, set the `Shell path` to `C:\Program Files\Git\bin\bash.exe`

The above path is for a default git installation location. You may need to adjust it in case you have installed git at a
different location.

#### 4. Open a built-in terminal

`View` > `Tool Windows` > `Terminal`

#### 5. Launch TCR

> ***Reminder***: the command below should be run from [Kata-Slow-Code/python](.) directory

From the built-in terminal:

```shell
./tcrw
```

Refer to [Using TCR](#using-tcr) section for additional details about TCR and available options.

<a name="running-the-kata-from-visual-studio-code"/></a>

### Running the kata from Visual Studio Code

#### 1. Prerequisite: install PHP Debug extension for Visual Studio Code

- Menu `View` > `Extensions`
- Search for `php`
- Select [PHP Debug](https://marketplace.visualstudio.com/items?itemName=xdebug.php-debug) and click `Install`

#### 2. Open the kata

Choose `Open Folder`, navigate to `Kata-Slow-Code` / `php` then click `Select Folder`.

#### 4. Select python interpreter

Click on `Select Interpreter` in the bottom bar, and choose the interpreter located
under `./venv` (should be the one tagged as `Recommended`)

#### 5. Run the tests

From the `Test Explorer` window (menu `View` > `Testing`):

- `Configure Python Tests` > `pytest` > `tests`
- Click on the `Play` button to run the tests.

<a name="running-the-kata-from-visual-studio-code-with-tcr"/></a>

### Running the kata from Visual Studio Code with TCR

TCR is provided as a command line utility running in a terminal.
You can run it from Visual Studio Code directly, through leveraging on its built-in terminal.

#### 1. Prerequisite: install Python extension for Visual Studio Code

- Menu `View` > `Extensions`
- Search for `python`
- Select [Python](https://marketplace.visualstudio.com/items?itemName=ms-python.python) and click `Install`

#### 2. Open the kata

Open Visual Studio Code, choose `Open Folder`, navigate to `Kata-Slow-Code` / `python`
then click `Select Folder`.

#### 3. Turn off auto-save

TCR is constantly watching the filesystem for changes.
For this reason you need to make sure that Visual Studio Code's auto-save is turned off
in order for TCR to behave as expected.

`File` > `Preferences` > `Settings`

In `Text Editor` > `Files` section, make sure that `Auto Save` setting is set to `off`

#### 4. Configure the built-in terminal to run git bash

> ***Windows Only***
>
> Skip this step if you're on macOS or Linux

Visual Studio Code for Windows is usually set up to run PowerShell by default in its built-in terminal.
TCR does not run in PowerShell.

`File` > `Preferences` > `Settings`

In `Features` > `Terminal` section, set `External: Windows Exec`
to `C:\Program Files\Git\bin\bash.exe`

The above path is for a default git installation location. You may need to adjust it in case you have installed git at a
different location.

#### 5. Open a built-in terminal

`Terminal` > `New Terminal`

#### 6. Start python virtual environment

> ***Reminder***: the command below should be run from [Kata-Slow-Code/python](.) directory

From the built-in terminal:

```shell
./start_python_venv.sh
```

#### 6. Launch TCR

> ***Reminder***: the command below should be run from [Kata-Slow-Code/python](.) directory

From the built-in terminal:

```shell
./tcrw
```

Refer to [Using TCR](#using-tcr) section for additional details about TCR and available options.

<a name="using-tcr"/></a>

## Using TCR

### Cheat Sheet

Here are the main shortcuts available once TCR utility is running:

| Shortcut  | Description                                   |
|-----------|-----------------------------------------------|
| `?`       | List available options                        |
| `d` / `D` | Enter driver role (from main menu)            |
| `n` / `N` | Enter navigator role (from main menu)         |
| `q` / `Q` | Quit current role - Quit TCR (from main menu) |
| `t` / `T` | Query timer status (from driver role only)    |
| `p` / `P` | Toggle on/off git auto-push (from main menu)  |

### Additional Details

Refer to [TCR - Test && Commit || Revert](../tcr/TCR.md) page
for additional details and explanations about TCR utility.
