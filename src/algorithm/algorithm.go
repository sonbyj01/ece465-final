package algorithm

import "models"

func Select(p models.Point2D, q models.Point2D, S models.Graph2D) {
	for _, point := range S.Points {
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
}