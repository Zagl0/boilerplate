package main

import "testing"

func TestAppBehavior(t *testing.T) {
    if !shouldAppPass() {
        t.Errorf("App is configured to fail")
    }
}