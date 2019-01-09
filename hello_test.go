package main

import "testing"

func TestHello(t *testing.T)  {
  actual := helloworld_function()
  expected := "hello world! \n"
  
  if actual != expected {
    t.Errorf("Actual '%s' Expected '%s'", actual, expected)
  }
  
}