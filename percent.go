package percent

func Calculate(current int, all int) float64 {
  percent := ( float64(current) * float64(100) ) / float64(all)
  return percent
}
