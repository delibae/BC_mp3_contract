module project.com/BCM

go 1.16

require (
	project.com/BC v0.0.0
	project.com/mp3_bytes v0.0.0

)

replace (
	project.com/BC v0.0.0 => ./BC
	project.com/mp3_bytes v0.0.0 => ./mp3_bytes

)
