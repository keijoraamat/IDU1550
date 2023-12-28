    	initialVector := p.VectorTo(*rotationP)

    newX := initialVector.X()*math.Cos(angle) - initialVector.Y()*math.Sin(angle)
    newY := initialVector.X()*math.Sin(angle) + initialVector.Y()*math.Cos(angle)
    expectedX := rotationP.X() + newX
    expectedY := rotationP.Y() + newY

    	if math.Abs(p.X()-expectedX) > tolerance || math.Abs(p.Y()-expectedY) > tolerance {
    	t.Errorf("Rotate(%f) = {%f, %f}, want {%f, %f}", angle, p.X(), p.Y(), expectedX, expectedY)
    }
