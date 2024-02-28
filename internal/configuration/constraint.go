package configuration

type Constraint func(map[int]*InternalParameter) (bool, error)

func CreateContraint(condition condition, exexution execution) Constraint {
	return func(config map[int]*InternalParameter) (bool, error) {
		if condition.fulfilled(config) {
			exexution.execute(config)
		}
		return false, nil
	}
}
