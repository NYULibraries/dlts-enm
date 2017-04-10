package cache

// TODO:
//   * Figure out if there's a better place to put this
//   * Allow user to specify cache path
// Tried using os.TempDir(), but it was returning
// /var/folders/dh/48wd7vnj3xqd1w_f126tcnvh0000gn/T/, which was not as convenient.
var Cache = "/tmp"