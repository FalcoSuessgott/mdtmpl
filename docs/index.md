# mdtmpl
Tired of copy-pasting your example configurations and command outputs into your README?

`mdtmpl` is a dead-simple little CLI tool that runs instructions defined in [Markdown comments](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#hiding-content-with-comments), such as `<!--- {{ my markdown comment }} --->`.

## Example
Imagine the following `README.md.tmpl`, when invoked, `mdtmpl` will interpret and render the instructions defined within `<!---{{ }}--->` to the following:

=== "`README.md.tmpl`"

    ```md
    ### Example Configuration
    Here are all available configuration options:
    <!--- {{ file "config.yml" | code "yaml" }} --->

    ### List Docker Containers
    You should now see docker containers running:
    <!--- {{ exec "docker ps -a" | truncate | code "bash" }} --->
    ```

=== "`README.md`"
    ### Example Configuration
    Here are all available configuration options:
    <!--- {{ file "config.yml" | code "yaml" }} --->
    ```yaml
    auth:
        basic: true
    ```

    ### List Docker Containers
    You should now see docker containers running:
    <!--- {{ exec "docker ps -a" | truncate | code "bash" }} --->
    ```bash
    CONTAINER ID   IMAGE         COMMAND                  CREATED       STATUS                   PORTS                                       NAMES
    cf4f9cec8faa   registry:2    "/entrypoint.sh /etc…"   7 weeks ago   Up 10 seconds            0.0.0.0:5000->5000/tcp, :::5000->5000/tcp   registry
    006560ea14d9   hello-world   "/hello"                 7 weeks ago   Exited (0) 7 weeks ago                                               dreamy_feistel
    d9d050df8a0f   hello-world   "/hello"                 7 weeks ago   Exited (0) 7 weeks ago
    ```
