output "with_sleep" {
  value = provider::sleep::with("hello", "3s")
  #=> value = "hello" (with 3s sleep)
}
