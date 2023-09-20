package database
import "fmt"
func (db *appdbimpl) UpdateAccount(updatedInfo User) (User, error) {
	fmt.Println(updatedInfo.ID)
	fmt.Println(updatedInfo.Username)
	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`, updatedInfo.Username, updatedInfo.ID)
	if err != nil {
		return updatedInfo, err
	}
	fmt.Println(err)
	fmt.Println(updatedInfo.Username)
	affected, err := res.RowsAffected()
	fmt.Println(affected)
	if err != nil {
		return updatedInfo, err
	} else if affected == 0 {
		fmt.Println("hey")
		// If we didn't update any row, then the user didn't exist
		return updatedInfo, ErrUserDoesNotExist
	}
	return updatedInfo, nil
}
