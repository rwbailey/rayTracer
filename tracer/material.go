package tracer

import (
	"math"
)

type Material struct {
	Colour    Colour
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
	Pattern   Pattern
}

func NewMaterial() *Material {
	return &Material{
		Colour:    NewColour(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
		Pattern:   nil,
	}
}

func (m *Material) Lighting(obj Shape, light *PointLight, point, eyev, normalv Tuple, inShadow bool) Colour {

	if m.Pattern != nil {
		m.Colour = m.Pattern.ColourAtObject(obj, point)
	}

	var ambient Colour
	var diffuse Colour
	var specular Colour

	// Combine the surface colour with the light's colour
	effectiveColour := m.Colour.Product(light.Intensity)

	// Find the direction of the light source
	lightv := light.Position.Subtract(point).Normalise()

	// Compute the ambiant contribution
	ambient = effectiveColour.Multiply(m.Ambient)

	// lightDotNormal represents the cosine of the angle between the light vector and the normal vector.
	// A negative number means the light is on the other side of the surface.
	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal < 0 {
		diffuse = NewColour(0, 0, 0)
		specular = NewColour(0, 0, 0)
	} else {
		// Compute diffuse contribution
		diffuse = effectiveColour.Multiply(m.Diffuse * lightDotNormal)

		// reflectDoteye represents the cosine of the angle between the reflection vector and the eye
		// vector.
		// A negative number means the light reflects away from the eye.
		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)
		if reflectDotEye <= 0 {
			specular = NewColour(0, 0, 0)
		} else {
			// Compute the specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Multiply(m.Specular * factor)
		}
	}
	if inShadow {
		return ambient
	}
	return ambient.Add(diffuse).Add(specular)
}
