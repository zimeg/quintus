variable "domain" {
  description = "Timekeeper online"
  type        = string
  default     = "quintus.sh"
}

variable "image" {
  description = "Flaked snapshots"
  type        = string
}

variable "vhs" {
  description = "Storage of the image"
  type        = string
  default     = "quintus.vhs"
}
