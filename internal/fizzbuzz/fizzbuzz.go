package fizzbuzz

import "strconv"

func Generate(int1, int2, limit int, str1, str2 string) ([]string, error) {
	if err := validate(int1, int2, limit); err != nil {
		return nil, err
	}

	result := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		switch {
		case i%int1 == 0 && i%int2 == 0:
			result = append(result, str1+str2)
		case i%int1 == 0:
			result = append(result, str1)
		case i%int2 == 0:
			result = append(result, str2)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result, nil
}

func validate(int1, int2, limit int) error {
	if int1 <= 0 || int2 <= 0 {
		return ErrInvalidDivisor
	}
	if limit <=0 {
		return ErrInvalidLimit
	}
	return nil
}