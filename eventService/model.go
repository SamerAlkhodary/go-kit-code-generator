package eventService

type Event struct {
	Title string `json:"title"`

	Id int64 `json:"id"`

	Description string `json:"description"`

	StartDate string `json:"startDate"`

	EndDate string `json:"endDate"`

	Image string `json:"image"`

	Url string `json:"url"`

	Images []string `json:"images"`

	EventSort string `json:"eventSort"`

	Location Location `json:"location"`

	Organizer Organizer `json:"organizer"`

	Offer Offer `json:"offer"`
}

func MakeNewEvent(title string, id int64, description string, startDate string, endDate string, image string, url string, images []string, eventSort string, location Location, organizer Organizer, offer Offer) *Event {
	return &Event{

		Title: title,

		Id: id,

		Description: description,

		StartDate: startDate,

		EndDate: endDate,

		Image: image,

		Url: url,

		Images: images,

		EventSort: eventSort,

		Location: location,

		Organizer: organizer,

		Offer: offer,
	}
}
func (event *Event) GetTitle() string {
	return event.Title
}

func (event *Event) SetTitle(newTitle string) {
	event.Title = newTitle
}

func (event *Event) GetId() int64 {
	return event.Id
}

func (event *Event) SetId(newId int64) {
	event.Id = newId
}

func (event *Event) GetDescription() string {
	return event.Description
}

func (event *Event) SetDescription(newDescription string) {
	event.Description = newDescription
}

func (event *Event) GetStartDate() string {
	return event.StartDate
}

func (event *Event) SetStartDate(newStartDate string) {
	event.StartDate = newStartDate
}

func (event *Event) GetEndDate() string {
	return event.EndDate
}

func (event *Event) SetEndDate(newEndDate string) {
	event.EndDate = newEndDate
}

func (event *Event) GetImage() string {
	return event.Image
}

func (event *Event) SetImage(newImage string) {
	event.Image = newImage
}

func (event *Event) GetUrl() string {
	return event.Url
}

func (event *Event) SetUrl(newUrl string) {
	event.Url = newUrl
}

func (event *Event) GetImages() []string {
	return event.Images
}

func (event *Event) SetImages(newImages []string) {
	event.Images = newImages
}

func (event *Event) GetEventSort() string {
	return event.EventSort
}

func (event *Event) SetEventSort(newEventSort string) {
	event.EventSort = newEventSort
}

func (event *Event) GetLocation() Location {
	return event.Location
}

func (event *Event) SetLocation(newLocation Location) {
	event.Location = newLocation
}

func (event *Event) GetOrganizer() Organizer {
	return event.Organizer
}

func (event *Event) SetOrganizer(newOrganizer Organizer) {
	event.Organizer = newOrganizer
}

func (event *Event) GetOffer() Offer {
	return event.Offer
}

func (event *Event) SetOffer(newOffer Offer) {
	event.Offer = newOffer
}

type Location struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Url string `json:"url"`

	Telephone string `json:"telephone"`

	Address Address `json:"address"`

	Geo Geo `json:"geo"`

	Id int64 `json:"id"`
}

func MakeNewLocation(name string, description string, url string, telephone string, address Address, geo Geo, id int64) *Location {
	return &Location{

		Name: name,

		Description: description,

		Url: url,

		Telephone: telephone,

		Address: address,

		Geo: geo,

		Id: id,
	}
}
func (location *Location) GetName() string {
	return location.Name
}

func (location *Location) SetName(newName string) {
	location.Name = newName
}

func (location *Location) GetDescription() string {
	return location.Description
}

func (location *Location) SetDescription(newDescription string) {
	location.Description = newDescription
}

func (location *Location) GetUrl() string {
	return location.Url
}

func (location *Location) SetUrl(newUrl string) {
	location.Url = newUrl
}

func (location *Location) GetTelephone() string {
	return location.Telephone
}

func (location *Location) SetTelephone(newTelephone string) {
	location.Telephone = newTelephone
}

func (location *Location) GetAddress() Address {
	return location.Address
}

func (location *Location) SetAddress(newAddress Address) {
	location.Address = newAddress
}

func (location *Location) GetGeo() Geo {
	return location.Geo
}

func (location *Location) SetGeo(newGeo Geo) {
	location.Geo = newGeo
}

func (location *Location) GetId() int64 {
	return location.Id
}

func (location *Location) SetId(newId int64) {
	location.Id = newId
}

type Organizer struct {
	Name string `json:"name"`

	Logo string `json:"logo"`

	Url string `json:"url"`

	Email string `json:"email"`

	Telephone string `json:"telephone"`
}

func MakeNewOrganizer(name string, logo string, url string, email string, telephone string) *Organizer {
	return &Organizer{

		Name: name,

		Logo: logo,

		Url: url,

		Email: email,

		Telephone: telephone,
	}
}
func (organizer *Organizer) GetName() string {
	return organizer.Name
}

func (organizer *Organizer) SetName(newName string) {
	organizer.Name = newName
}

func (organizer *Organizer) GetLogo() string {
	return organizer.Logo
}

func (organizer *Organizer) SetLogo(newLogo string) {
	organizer.Logo = newLogo
}

func (organizer *Organizer) GetUrl() string {
	return organizer.Url
}

func (organizer *Organizer) SetUrl(newUrl string) {
	organizer.Url = newUrl
}

func (organizer *Organizer) GetEmail() string {
	return organizer.Email
}

func (organizer *Organizer) SetEmail(newEmail string) {
	organizer.Email = newEmail
}

func (organizer *Organizer) GetTelephone() string {
	return organizer.Telephone
}

func (organizer *Organizer) SetTelephone(newTelephone string) {
	organizer.Telephone = newTelephone
}

type Geo struct {
	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`
}

func MakeNewGeo(latitude float64, longitude float64) *Geo {
	return &Geo{

		Latitude: latitude,

		Longitude: longitude,
	}
}
func (geo *Geo) GetLatitude() float64 {
	return geo.Latitude
}

func (geo *Geo) SetLatitude(newLatitude float64) {
	geo.Latitude = newLatitude
}

func (geo *Geo) GetLongitude() float64 {
	return geo.Longitude
}

func (geo *Geo) SetLongitude(newLongitude float64) {
	geo.Longitude = newLongitude
}

type Address struct {
	StreetAddress string `json:"streetAddress"`

	AddressLocality string `json:"addressLocality"`

	AddressCounty string `json:"addressCounty"`

	PodtalCode string `json:"podtalCode"`
}

func MakeNewAddress(streetAddress string, addressLocality string, addressCounty string, podtalCode string) *Address {
	return &Address{

		StreetAddress: streetAddress,

		AddressLocality: addressLocality,

		AddressCounty: addressCounty,

		PodtalCode: podtalCode,
	}
}
func (address *Address) GetStreetAddress() string {
	return address.StreetAddress
}

func (address *Address) SetStreetAddress(newStreetAddress string) {
	address.StreetAddress = newStreetAddress
}

func (address *Address) GetAddressLocality() string {
	return address.AddressLocality
}

func (address *Address) SetAddressLocality(newAddressLocality string) {
	address.AddressLocality = newAddressLocality
}

func (address *Address) GetAddressCounty() string {
	return address.AddressCounty
}

func (address *Address) SetAddressCounty(newAddressCounty string) {
	address.AddressCounty = newAddressCounty
}

func (address *Address) GetPodtalCode() string {
	return address.PodtalCode
}

func (address *Address) SetPodtalCode(newPodtalCode string) {
	address.PodtalCode = newPodtalCode
}

type Offer struct {
	Price string `json:"price"`

	Url string `json:"url"`
}

func MakeNewOffer(price string, url string) *Offer {
	return &Offer{

		Price: price,

		Url: url,
	}
}
func (offer *Offer) GetPrice() string {
	return offer.Price
}

func (offer *Offer) SetPrice(newPrice string) {
	offer.Price = newPrice
}

func (offer *Offer) GetUrl() string {
	return offer.Url
}

func (offer *Offer) SetUrl(newUrl string) {
	offer.Url = newUrl
}
