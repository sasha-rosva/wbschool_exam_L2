package pattern

/*
Суть паттерна builder в упрощении логики конструирования какого-то сложного объекта. Плюсы: позволяет создавать продукты пошагово, позволяет использовать один и тот же код для 
создания различных продуктов,изолирует сложный код сборки продукта от его основной бизнес-логики. Минусы: усложняет код программы из-за введения дополнительных структур.
*/
import "fmt"
type russianBuilder struct {
	sideDish string
	soap   string
	mainDish      string
}

func newRussianBuilder() *russianBuilder {
	return &russianBuilder{}
}

func (b *russianBuilder) cookSideDish() {
	b.sideDish = "Заливное из судака"
}

func (b *russianBuilder) cookSoap() {
	b.soap = "Борщ"
}

func (b *russianBuilder) cookMainDish() {
	b.mainDish = "Бефстроганов из говядины"
}

func (b *russianBuilder) getLaunch() launch {
	return launch{
		sideDish:   b.sideDish,
		soap: b.soap,
		mainDish:      b.mainDish,
	}
}
type mexicanBuilder struct {
	sideDish string
	soap   string
	mainDish      string
}

func newMexicanBuilder() *mexicanBuilder {
	return &mexicanBuilder{}
}

func (b *mexicanBuilder) cookSideDish() {
	b.sideDish = "Кесадилья"
}

func (b *mexicanBuilder) cookSoap() {
	b.soap = "Чили кон карне"
}

func (b *mexicanBuilder) cookMainDish() {
	b.mainDish = "Фахитос"
}

func (b *mexicanBuilder) getLaunch() launch {
	return launch{
		sideDish:   b.sideDish,
		soap: b.soap,
		mainDish:      b.mainDish,
	}
}
type launch struct {
	sideDish string
	soap   string
	mainDish      string
}
type iBuilder interface {
	cookSideDish()
	cookSoap()
	cookMainDish()
	getLaunch() launch
}

func getBuilder(builderType string) iBuilder {
	if builderType == "russian" {
		return &russianBuilder{}
	}

	if builderType == "mexican" {
		return &mexicanBuilder{}
	}
	return nil
}
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) cookLaunch() launch {
	d.builder.cookSideDish()
	d.builder.cookSoap()
	d.builder.cookMainDish()
	return d.builder.getLaunch()
}
func main() {
	russianBuilder := getBuilder("russian")
	mexicanBuilder := getBuilder("mexican")

	director := newDirector(russianBuilder)
	russianLaunch := director.cookLaunch()

	fmt.Printf("Русский ланч (Закуска): %s\n", russianLaunch.sideDish)
	fmt.Printf("Русский ланч (Суп): %s\n", russianLaunch.soap)
	fmt.Printf("Русский ланч (Основное блюдо): %s\n", russianLaunch.mainDish)

	director.setBuilder(mexicanBuilder)
	mexicanLaunch := director.cookLaunch()

	fmt.Printf("\nМексиканский ланч (Закуска): %s\n", mexicanLaunch.sideDish)
	fmt.Printf("Мексиканский ланч (Суп): %s\n", mexicanLaunch.soap)
	fmt.Printf("Мексиканский ланч (Основное блюдо): %s\n", mexicanLaunch.mainDish)

}
