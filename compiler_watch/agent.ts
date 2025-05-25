import { log } from "./log.js";
import ObjC from "frida-objc-bridge";

log("Hello from Frida:", Frida.version);
log(ObjC.available);
