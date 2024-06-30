# Dify Unsandboxed Sandbox

🌏 [**English**](./README.md)
🌏 [**日本語**](./README.ja.md)

> [!CAUTION]
>
> - **Using this container image will make your container host extremely vulnerable to attacks from malicious code.**
> - **Do not use this container image unless you understand what you are trying to do.**
> - **Do not use in production environment.**
> - **Do not use in Dify environment shared with others.**
> - **Use only in an environment that only you can use, for experiments and tests only.**

## Overview

This is a container image to replace langgenius/dify-sandbox, developed with the aim of removing all constraints of Dify's sandbox environment.

Dify's sandbox environment imposes various constraints on the code running in it to protect the environment from malicious (or unintentionally malicious) code.

This container image removes those constraints and makes it possible to run any code in the sandbox.

| Blocks | Languages | Tested |
| --- | --- | :---: |
| **Code** | Python | ✅ |
| **Code** | Node.js | - |
| **Template** | Python | ✅ |

## Disclaimer

**Using this container image means that your environment is extremely vulnerable to malicious code.**

- 🚨 **The code run by the user is executed by the privileged user in the container.**
- 🚨 **The code run by the user can access all files and processes in the container.**
- 🚨 **The code run by the user can access any network that the container can connect to.**
- 🚨 **If there is a container escape vulnerability in the container environment, the impact will affect not only the container but also the host and the entire environment to which it is connected.**

**Do not use this container image unless you understand what you are trying to do.**

**To confirm how to use it with full knowledge of the risks, please refer to [the usage page](./docs/usage.md).**
