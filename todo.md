# TODO

- [ ] Formatting of gh command, spinner should be ignored

- [ ] Check if file has execution permissions
- [ ] Add ability to register keybindings
- [ ] improve layout renderer, https://github.com/rivo/tview
- [ ] use cli lib for adding commands https://github.com/cristalhq/acmd
- [ ] Use Go template for adding secrets

# IN-PROGRESS

# DONE

- [x] Ensure all modules running are killed on exit
- [x] Rethink usage of space at top of dashboard
- [x] Use string for width so you can add "full" for fullwidth
- [x] bug: app craches if names are more than one word
- [x] Add loading state to modules
- [x] build tabs
- [x] Only run modules for current tab
- [x] toml config file
- [x] render a component for each module
- [x] Prevent rerunning of modules
- [x] Prevent mutation of modules?
- [x] create modulebox with title
- [x] Get true width and height of output
- [x] Add layout engine
- [x] Add CLICOLOR_FORCE
- [x] Add keybinding for opening config file
- [x] Create basic template for config
- [x] use xdg paths for config
- [x] Update struct for config v2
- [x] Validate config file and all modules
- [x] terminal doesn't clear after quitting and having opened config file
- [x] Better logger
- [x] Fix opening config action https://gist.github.com/bashbunni/e3306e8633512d8134012028288212db
- [x] Create module error
- [x] cd to config folder for running module scripts
- [x] Add status bar at bottom
- [x] Cancel module when changing tab
