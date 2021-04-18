package main

type (
	MovieLister struct {
		finder MoviesFinder
	}
	MoviesFinder interface {
		FindAll() []Movie
	}

	Movie struct {
		Director string
	}
)

func (ml *MovieLister) MoviesDirectedBy(director string) []Movie {
	allMovies := ml.finder.FindAll()
	result := make([]Movie, 0, len(allMovies))

	for _, m := range allMovies {
		if director == m.Director {
			result = append(result, m)
		}
	}

	return result
}
