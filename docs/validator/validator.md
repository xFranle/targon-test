# Running a Validator

**Validator Prerequisites**

- Basic compute resources (any CPU server)

## Validator Installation

**1. Set Up the Validator CPU Server**

Deploy on an AMD SEV-SNP enabled server. We recommend
[Latitude](https://www.latitude.sh/) m4.metal.medium boxes running Ubuntu
26.04.

**2. Environment Setup**

Create a `.env` file in the project directory:

```bash
### Required
HOTKEY_PHRASE="your hotkey phrase"
MONGO_USERNAME=
MONGO_PASSWORD=
VALIDATOR_IP="ipv4 address of the vali"

### Optional

# Sets a multiplier on all http connection timeouts
TIMEOUT_MULT=1

## Sends discord notifications for things like setting weights
# DISCORD_URL=

### Dev Variables (don't set unless needed)

## Netuid the validator is running on. Useful for testnet
# NETUID=4

## Chain to connect to
# CHAIN_ENDPOINT=wss://entrypoint-finney.opentensor.ai:443
```

> **Note:** Both Mongo environment keys are values **you define** and can be
> anything you want. We suggest using the output of
> `tr -dc A-Za-z0-9 </dev/urandom | head -c 24; echo` — run it twice, and use
> one result as the username and the other as the password. These **do not**
> need to be double-quoted in the `.env` file.

**3. Pull the Validator VM**

Run the `tvm/install` binary on the CPU server with the following arguments:

- `--hotkey-phrase`: Your validator hotkey phrase
- `--node-type`: `vali-cpu`
- `--submit`: Actually submit and download the image
- `--service-url`: `http://tvm.targon.com`
- `--vm-download-dir`: Location to download the VM to (e.g. `/vms/manifold`)

**4. Start the VM**

Go to the VM download directory and enter the newly created (unzipped) folder,
then start the VM:

```bash
sudo launch_vm.sh
```

To stop the VM, find the QEMU process ID and kill it:

```bash
ps -aux | grep qemu
sudo kill -9 [pid]
```

where `[pid]` is the second column of the output row that starts with `root`.

**5. Start the Validator**

To initialize the validator, first install `sn4` **on the server hosting the
VM** (see the CLI docs for installation), then run:

```bash
sn4 vali init [path to .env file]
```

You will be prompted to confirm the environment file before it is applied. This
command must be run on the server where the VM is running; otherwise, pass
`--ip` with the IP address of that server. Once confirmed, the validator starts
with the provided environment variables.

## Validator Monitoring and Maintenance

**Logs**

First, list the running containers:

```bash
sn4 vali containers
```

Then view the logs of any container:

```bash
sn4 vali logs --container [container name]
```

**Updates**

To update the validator running in the VM, either re-pull and re-initialize the
VM, or run:

```bash
sn4 vali update
```

