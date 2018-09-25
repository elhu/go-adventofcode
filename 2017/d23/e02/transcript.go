a, b, d, d, e, f, g, h := 0, 0, 0, 0, 0, 0, 0, 0

a = 1
b = 67
if a != 0 {
  b *= 100
  b += 100000
  c = b
  c += 17000
}
for {
  f = 1
  d = 2
  for {
    e = 2
    for {
      g = d
      g *= e
      g -= b
      if g == 0 {
        f = 0
      }
      e += 1
      g = e
      g -= b
      if g == 0 {
        break
      }
    }
    d += 1
    g = d
    g -= b
    if g == 0 {
      break
    }
  }
  if f == 0 {
    h += 1
  }
  g = b
  g -= c
  if g == 0 {
    break
  }
  b += 17
}
