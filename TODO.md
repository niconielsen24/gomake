REDO build system:

general idea:

- langs | templates  -> 2 tabs to switch on home

langs : select a language and show available options
        maybe a good idea to divide options into categories as to not
        select multiple conflicting options

templetas : should be just a predefined build that jumps straight to 
            acceptance stage

- preview -> describes the current build at this stage
             could be a right side screen all the time,
             or just a screen at the final stage

- building -> should be an interface which provides the Build() methods
              maybe also  : BuildWithModules(string...)
                            BuildWithLib(string...)
              also building is not a really good name, as we are no really 
              building anything
