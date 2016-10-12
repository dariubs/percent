package percent

func Percent(pcent int, all int) float64{
  percent := ((float64(all) * float64(pcent)) / float64(100))
  return percent
}

func PercentOf(current int, all int) float64 {
  percent := ( float64(current) * float64(100) ) / float64(all)
  return percent
}

func Change(before int, after int) float64{
  diff := float64(after) - float64(before)
  realDiff := diff / float64(before)
  percentDiff := 100 * realDiff

  return percentDiff
}
