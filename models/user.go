package models

type User struct {
}

func (user *User) SerializeJson() ([]byte, error) {
    data, err := json.MarshalIndent(*user, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return nil, err
    }

    return data, nil
}

func (user *User) DeserializeJson(obj []byte) error {
    err := json.Unmarshal(obj, user)

    if err != nil {
        return err
    }

    return nil
}

func (user *User) Expand(dbmodels.User) error {
    return nil
}

func (user *User) Collapse() (*dbmodels.User, error) {
    return nil, nil
}
