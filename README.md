# lawnchair

A simple launcher that will start up any process, and supply any needed cli args.  

## Sample config

```yaml
programs:
    game:
        args:
            - Project Test Game
        description: Test Game built in Xcode using Swift
        name: "game"
        path: /Users/me/bin/mygame
    game2:
        args:
            - sample1
            - --config
            - ./config.json
        description: Game2 program
        name: "game2"
        path: /Users/me/bin/lawnchair
    worker:
        args:
            - sample2
        description: Worker program
        env_vars:
            WORKER_HOST: localhost
            WORKER_ID: worker1
            WORKER_PORT: "5000"
        name: "worker"
        path: /Users/me/bin/lawnchair
```
