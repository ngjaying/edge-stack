## Steps

1. Build docker image

```shell
edge-stack$ sudo docker image build -t sim:1 -f docker/Dockerfile .
```

2. Run docker compose sim

```shell
edge-stack/developer-scripts$ docker-compose -f "docker-compose_sim.yml" -p "sim_stack" up -d
```

3. Verify the password loading log. Make sure password is set and the password value is not printed

## Thinking

1. The linux permission
    1. Folder "x": can mount into the folder
    2. Folder "w": can delete or add files in that folder
    3. File "w": can modify the file, but cannot delete if folder is not writable.
2. Docker volume permission: lazy set
    1. Set by the creation docker image. Can run `chmod` in the dockerfile
    2. Set in the host system by finding the physical volume path.