# sn4 CLI

## Installation

Install `sn4` from the Targon APT repository:

```bash
sudo curl -fsSL https://releases.targon.com/targon.gpg -o /usr/share/keyrings/targon.gpg
echo "deb [signed-by=/usr/share/keyrings/targon.gpg] https://releases.targon.com/apt stable main" | sudo tee /etc/apt/sources.list.d/targon.list
sudo apt update
sudo apt install sn4
```

Updates then arrive through `apt upgrade` like any other package. To uninstall, run `sudo apt remove sn4`.

### Building from source

Alternatively, build and install from this repository:

```bash
make install-cli
```

This installs the binary to `/usr/local/bin/sn4`. Depending on your permissions, you may need to run `sudo make install-cli`. To install to a different location, override `PREFIX` (e.g. `make install-cli PREFIX=$HOME/.local` installs to `~/.local/bin/sn4` — make sure that `bin` directory is in your `PATH`). To uninstall, run `make uninstall-cli` with the same `PREFIX`.

## Quick Start

1. **Install the CLI:**
  ```bash
   sudo apt install sn4
  ```
2. **Run first-time setup:**
  ```bash
   sn4
  ```
   On first run, the CLI prompts for your hotkey phrase and creates a configuration file at `~/.config/.targon.json`.

## Usage

```bash
sn4 [command] [flags]
```

## Commands

### `sn4 attest`

Validator-only command. Manually attest a miner or IP address for attestation.

**Usage:**

```bash
sn4 attest [flags]
```

**Flags:**


| Flag    | Type   | Description                               |
| ------- | ------ | ----------------------------------------- |
| `--uid` | int    | UID of the miner to attest                |
| `--ip`  | string | Specific IP address for off-chain testing |


**Examples:**

```bash
# Attest a specific UID
sn4 attest --uid 123

# Attest a specific IP address
sn4 attest --ip 192.168.1.100
```

### `sn4 config`

Update configuration settings.

**Usage:**

```bash
sn4 config [flags]
```

**Flags:**


| Flag                        | Type   | Description                 |
| --------------------------- | ------ | --------------------------- |
| `--miner.hotkey-phrase`     | string | New miner hotkey phrase     |
| `--validator.hotkey-phrase` | string | New validator hotkey phrase |


**Examples:**

```bash
# Update the miner hotkey phrase
sn4 config --miner.hotkey-phrase "your hotkey phrase here"

# Update the validator hotkey phrase
sn4 config --validator.hotkey-phrase "your hotkey phrase here"
```

### `sn4 get`

Fetch data from MongoDB or the blockchain and display it in various formats.

**Usage:**

```bash
sn4 get [command]
```

**Subcommands:**

#### `sn4 get errors`

Get attestation errors for a specific UID.

**Usage:**

```bash
sn4 get errors [flags]
```

**Flags:**


| Flag    | Type | Description                         |
| ------- | ---- | ----------------------------------- |
| `--uid` | int  | UID to fetch attestation errors for |


**Example:**

```bash
sn4 get errors --uid 123
```

## Configuration

The CLI stores its settings in a JSON file at `~/.config/.targon.json`, created during first-time setup. It holds your hotkey phrase and other settings.

**Configuration file structure:**

```json
{
  "HOTKEY_PHRASE": "your-hotkey-phrase-here"
}
```

