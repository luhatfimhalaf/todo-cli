package tasks

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// SaveToFile menyimpan TaskList ke file JSON
   func (tl *TaskList) SaveToFile(filename string) error {
       data, err := json.MarshalIndent(tl, "", "  ")
       if err != nil {
           return err
       }
       return ioutil.WriteFile(filename, data, 0644)
   }

   // LoadFromFile memuat TaskList dari file JSON
   func LoadFromFile(filename string) (*TaskList, error) {
       file, err := ioutil.ReadFile(filename)
       if err != nil {
           if os.IsNotExist(err) {
               return &TaskList{Tasks: []Task{}}, nil
           }
           return nil, err
       }
       var tl TaskList
       err = json.Unmarshal(file, &tl)
       if err != nil {
           return nil, err
       }
       return &tl, nil
   }