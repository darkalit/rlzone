package pagination

import "math"

type Pagination struct {
	TotalCount int64
	TotalPages int64
	Size       int64
	Page       int64
	HasMore    bool
	HasLess    bool
}

const (
	defaultSize = 15
	defaultPage = 1
)

func SanitizePages(page *int, size *int) {
	if *page == 0 {
		*page = defaultPage
	}

	if *size == 0 {
		*size = defaultSize
	}
}

func GetOffsetAndLimit(page int, size int) (offset int, limit int) {
	return (page - 1) * size, size
}

func GetPagination(page int64, size int64, totalCount int64) Pagination {
	totalPages := int64(math.Ceil(float64(totalCount) / float64(size)))

	return Pagination{
		TotalCount: totalCount,
		TotalPages: totalPages,
		Size:       size,
		Page:       page,
		HasMore:    page < totalPages,
		HasLess:    page > 1,
	}
}
