import createPlugin from "@extism/extism";
import { DurableObject } from "cloudflare:workers";

export class Enclave extends DurableObject {
  constructor(ctx, env) {
    super(ctx, env);
    this.state = ctx;
    this.env = env;
    this.initializePlugins();
  }

  // initializePlugins initializes the signer plugin
  async initializePlugins() {
    try {
      // 1. Initialize the enclave plugin
      this.enclavePlugin = await createPlugin(
        "https://cdn.sonr.io/bin/enclave.wasm",
        {
          useWasi: true,
          config: {},
        },
      );
      return true;
    } catch (error) {
      this.addLog(`Failed to initialize plugin: ${error.message}`);
      console.error("Failed to initialize plugin:", error);
      return false;
    }
  }
}
