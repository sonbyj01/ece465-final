package algorithm

import "models"

func Select(S models.Graph2D) (models.Point2D, models.Point2D) {
	var p,q models.Point2D
	for index, point := range S.Points {
		if index == 0 {
			p = point
			continue
		}

		if index == 1 {
			q = point
			continue
		}

		// maximum y-value, then maximum x-value if y values are same
		if p.YValue < point.YValue {
			p = point
		} else if p.YValue == point.YValue {
			if p.XValue < point.XValue {
				p = point
			}
		}
		// minimum y-value, then minimum x-value if y values are same
		if q.YValue > point.YValue {
			q = point
		} else if q.YValue == point.YValue {
			if q.XValue > point.XValue {
				q = point
			}
		}
	}
	return p, q
}