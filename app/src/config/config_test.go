package config

import (
    "testing"
    "model"
)

func TestConfig(t *testing.T) {
    _, tags, _ := Config()

    tags.Create("tag #1")
    existing := (tags.GetAll().Value).(*[]model.Tag)

    if(len(*existing) != 1) {
        t.Log("One tag should be added!")
        t.Fail()
    }

    if((*existing)[0].Name != "tag #1") {
        t.Log("Incorrect tag name!")
        t.Fail()
    }
}