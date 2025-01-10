package pom

import (
	"encoding/xml"
	"io"
	"strings"
)

// Official Pom Schema https://maven.apache.org/xsd/maven-4.0.0.xsd
type Model struct {
	XMLName        xml.Name `xml:"project"`
	// Declares to which version of project descriptor this POM conforms.
	ModelVersion string `xml:"modelVersion,omitempty"`
	// The location of the parent project, if one exists. Values from the parent project will be the default for this project if they are left unspecified. The location is given as a group ID, artifact ID and version.
	Parent *Parent `xml:"parent,omitempty"`
	// A universally unique identifier for a project. It is normal to use a fully-qualified package name to distinguish it from other projects with a similar name (eg. <groupId>org.apache.maven</groupId>).
	GroupId string `xml:"groupId,omitempty"`
	// The identifier for this artifact that is unique within the group given by the group ID. An artifact is something that is either produced or used by a project. Examples of artifacts produced by Maven for a project include: JARs, source and binary distributions, and WARs.
	ArtifactId string `xml:"artifactId,omitempty"`
	// The current version of the artifact produced by this project.
	Version string `xml:"version,omitempty"`
	// The type of artifact this project produces, for example <code>jar</code> <code>war</code> <code>ear</code> <code>pom</code>. Plugins can create their own packaging, and therefore their own packaging types, so this list does not contain all possible types.
	Packaging string `xml:"packaging,omitempty"`
	// The full name of the project.
	Name string `xml:"name,omitempty"`
	// A detailed description of the project, used by Maven whenever it needs to describe the project, such as on the web site. While this element can be specified as CDATA to enable the use of HTML tags within the description, it is discouraged to allow plain text representation. If you need to modify the index page of the generated web site, you are able to specify your own instead of adjusting this text.
	Description string `xml:"description,omitempty"`
	// The URL to the project's homepage.
	//
	// Default value is: parent value [+ path adjustment] + (artifactId or project.directory property), or just parent value if project's <code>child.project.url.inherit.append.path="false"</code>
	Url string `xml:"url,omitempty"`
	// The year of the project's inception, specified with 4 digits. This value is used when generating copyright notices as well as being informational.
	InceptionYear string `xml:"inceptionYear,omitempty"`
	// This element describes various attributes of the organization to which the project belongs. These attributes are utilized when documentation is created (for copyright notices and links).
	Organization *Organization `xml:"organization,omitempty"`
	// This element describes all of the licenses for this project. Each license is described by a <code>license</code> element, which is then described by additional elements. Projects should only list the license(s) that applies to the project and not the licenses that apply to dependencies. If multiple licenses are listed, it is assumed that the user can select any of them, not that they must accept all.
	Licenses *Licenses `xml:"licenses,omitempty"`
	// Describes the committers of a project.
	Developers *Developers `xml:"developers,omitempty"`
	// Describes the contributors to a project that are not yet committers.
	Contributors *Contributors `xml:"contributors,omitempty"`
	// Contains information about a project's mailing lists.
	MailingLists *MailingLists `xml:"mailingLists,omitempty"`
	// Describes the prerequisites in the build environment for this project.
	Prerequisites *Prerequisites `xml:"prerequisites,omitempty"`
	// The modules (sometimes called subprojects) to build as a part of this project. Each module listed is a relative path to the directory containing the module. To be consistent with the way default urls are calculated from parent, it is recommended to have module names match artifact ids.
	Modules *Modules `xml:"modules,omitempty"`
	// Specification for the SCM used by the project, such as CVS, Subversion, etc.
	Scm *Scm `xml:"scm,omitempty"`
	// The project's issue management system information.
	IssueManagement *IssueManagement `xml:"issueManagement,omitempty"`
	// The project's continuous integration information.
	CiManagement *CiManagement `xml:"ciManagement,omitempty"`
	// Distribution information for a project that enables deployment of the site and artifacts to remote web servers and repositories respectively.
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	// Properties that can be used throughout the POM as a substitution, and are used as filters in resources if enabled. The format is <code>&lt;name&gt;value&lt;/name&gt;</code>.
	Properties *Properties `xml:"properties,omitempty"`
	// Default dependency information for projects that inherit from this one. The dependencies in this section are not immediately resolved. Instead, when a POM derived from this one declares a dependency described by a matching groupId and artifactId, the version and other values from this section are used for that dependency if they were not already specified.
	DependencyManagement *DependencyManagement `xml:"dependencyManagement,omitempty"`
	// This element describes all of the dependencies associated with a project. These dependencies are used to construct a classpath for your project during the build process. They are automatically downloaded from the repositories defined in this project. See [the dependency mechanism] for more information.
	//
	// [the dependency mechanism]: https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
	// The lists of the remote repositories for discovering dependencies and extensions.
	Repositories *Repositories `xml:"repositories,omitempty"`
	// The lists of the remote repositories for discovering plugins for builds and reports.
	PluginRepositories *PluginRepositories `xml:"pluginRepositories,omitempty"`
	// Information required to build the project.
	Build *Build `xml:"build,omitempty"`
	// Deprecated: Now ignored by Maven.
	Reports *Reports `xml:"reports,omitempty"`
	// This element includes the specification of report plugins to use to generate the reports on the Maven-generated site. These reports will be run when a user executes <code>mvn site</code>. All of the reports will be included in the navigation bar for browsing.
	Reporting *Reporting `xml:"reporting,omitempty"`
	// A listing of project-local build profiles which will modify the build process when activated.
	Profiles *Profiles `xml:"profiles,omitempty"`
}

type Prerequisites struct {
	Comment string `xml:",comment"`
	// For a plugin project (packaging is <code>maven-plugin</code>), the minimum version of Maven required to use the resulting plugin.
	Maven string `xml:"maven,omitempty"`
}

type Modules struct {
	Comment string   `xml:",comment"`
	Module  []string `xml:"module,omitempty"`
}

type Licenses struct {
	Comment string    `xml:",comment"`
	License []License `xml:"license,omitempty"`
}

type License struct {
	Comment      string `xml:",comment"`
	Name         string `xml:"name,omitempty"`
	Url          string `xml:"url,omitempty"`
	Distribution string `xml:"distribution,omitempty"`
	Comments     string `xml:"comments,omitempty"`
}

type CiManagement struct {
	Comment   string     `xml:",comment"`
	System    string     `xml:"system,omitempty"`
	Url       string     `xml:"url,omitempty"`
	Notifiers *Notifiers `xml:"notifiers,omitempty"`
}

type Notifiers struct {
	Comment  string     `xml:",comment"`
	Notifier []Notifier `xml:"notifier,omitempty"`
}

type Notifier struct {
	Comment       string `xml:",comment"`
	Type          string `xml:"type,omitempty"`
	SendOnError   bool   `xml:"sendOnError,omitempty"`
	SendOnFailure bool   `xml:"sendOnFailure,omitempty"`
	SendOnSuccess bool   `xml:"sendOnSuccess,omitempty"`
	SendOnWarning bool   `xml:"sendOnWarning,omitempty"`
	Address       string `xml:"address,omitempty"`
	Configuration *DOM   `xml:"configuration,omitempty"`
}

type Scm struct {
	Comment string `xml:",comment"`
	/*
		The source control management system URL that describes the repository and how to connect to the repository.
		For more information, see the [URL format] and [list of supported SCMs]. This connection is read-only.

		[URL format]: https://maven.apache.org/scm/scm-url-format.html
		[list of supported SCMs]: https://maven.apache.org/scm/scms-overview.html

	*/
	Connection          string `xml:"connection,omitempty"`
	DeveloperConnection string `xml:"developerConnection,omitempty"`
	Tag                 string `xml:"tag,omitempty"`
	Url                 string `xml:"url,omitempty"`
}

type IssueManagement struct {
	Comment string `xml:",comment"`
	System  string `xml:"system,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type DependencyManagement struct {
	Comment      string        `xml:",comment"`
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
}

type Dependency struct {
	Comment    string      `xml:",comment"`
	GroupId    string      `xml:"groupId,omitempty"`
	ArtifactId string      `xml:"artifactId,omitempty"`
	Version    string      `xml:"version,omitempty"`
	Type       string      `xml:"type,omitempty"`
	Classifier string      `xml:"classifier,omitempty"`
	Scope      string      `xml:"scope,omitempty"`
	SystemPath string      `xml:"systemPath,omitempty"`
	Exclusions *Exclusions `xml:"exclusions,omitempty"`
	Optional   string      `xml:"optional,omitempty"`
}

type Exclusions struct {
	Comment   string      `xml:",comment"`
	Exclusion []Exclusion `xml:"exclusion,omitempty"`
}

type Exclusion struct {
	Comment    string `xml:",comment"`
	ArtifactId string `xml:"artifactId,omitempty"`
	GroupId    string `xml:"groupId,omitempty"`
}

type Parent struct {
	Comment      string `xml:",comment"`
	GroupId      string `xml:"groupId,omitempty"`
	ArtifactId   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type Developers struct {
	Comment   string      `xml:",comment"`
	Developer []Developer `xml:"developer,omitempty"`
}

type Developer struct {
	Comment         string `xml:",comment"`
	Id              string `xml:"id,omitempty"`
	Name            string `xml:"name,omitempty"`
	Email           string `xml:"email,omitempty"`
	Url             string `xml:"url,omitempty"`
	Organization    string `xml:"organization,omitempty"`
	OrganizationUrl string `xml:"organizationUrl,omitempty"`
	Roles           *Roles `xml:"roles,omitempty"`
	Timezone        string `xml:"timezone:omitempty"`
	Properties      *DOM   `xml:"properties,omitempty"`
}

type Roles struct {
	Comment string   `xml:",comment"`
	Role    []string `xml:"role,omitempty"`
}

type MailingLists struct {
	Comment     string        `xml:",comment"`
	MailingList []MailingList `xml:"mailingList,omitempty"`
}

type MailingList struct {
	Comment       string         `xml:",comment"`
	Name          string         `xml:"name,omitempty"`
	Subscribe     string         `xml:"subscribe,omitempty"`
	Unsubscribe   string         `xml:"unsubscribe,omitempty"`
	Post          string         `xml:"post,omitempty"`
	Archive       string         `xml:"archive,omitempty"`
	OtherArchives []OtherArchive `xml:"otherArchives,omitempty"`
}

type OtherArchive struct {
	Comment      string `xml:",comment"`
	OtherArchive string `xml:"otherArchive,omitempty"`
}

type Contributors struct {
	Comment     string        `xml:",comment"`
	Contributor []Contributor `xml:"contributor,omitempty"`
}

type Contributor struct {
	Developer
}

type Organization struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type Dependencies struct {
	Comment    string       `xml:",comment"`
	Dependency []Dependency `xml:"dependency,omitempty"`
}

type DistributionManagement struct {
	Comment            string                `xml:",comment"`
	Repository         *DeploymentRepository `xml:"repository,omitempty"`
	SnapshotRepository *DeploymentRepository `xml:"snapshotRepository,omitempty"`
	Site               *Site                 `xml:"site,omitempty"`
	DownloadUrl        string                `xml:"downloadUrl,omitempty"`
	Reloction          *Relocation           `xml:"relocation,omitempty"`
	Status             string                `xml:"status,omitempty"`
}

type DeploymentRepository struct {
	Comment       string            `xml:",comment"`
	UniqueVersion bool              `xml:"uniqueVersion,omitempty"`
	Releases      *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots     *RepositoryPolicy `xml:"snapshots,omitempty"`
	Id            string            `xml:"id,omitempty"`
	Name          string            `xml:"name,omitempty"`
	Url           string            `xml:"url,omitempty"`
	Layout        string            `xml:"layout,omitempty"`
}

type Repositories struct {
	Comment    string       `xml:",comment"`
	Repository []Repository `xml:"repository,omitempty"`
}

type PluginRepositories struct {
	Comment    string       `xml:",comment"`
	Repository []Repository `xml:"repository,omitempty"`
}

type Repository struct {
	Comment   string            `xml:",comment"`
	Releases  *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots *RepositoryPolicy `xml:"snapshots,omitempty"`
	Id        string            `xml:"id,omitempty"`
	Name      string            `xml:"name,omitempty"`
	Url       string            `xml:"url,omitempty"`
	Layout    string            `xml:"layout,omitempty"`
}

// Download policy
type RepositoryPolicy struct {
	Comment        string `xml:",comment"`
	Enabled        string `xml:"enabled,omitempty"`
	UpdatePolicy   string `xml:"updatPolicy,omitempty"`
	ChecksumPolicy string `xml:"checksumPolicy,omitempty"`
}

type Site struct {
	Comment string `xml:",comment"`
	Id      string `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type Relocation struct {
	Comment    string `xml:",comment"`
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Message    string `xml:"message,omitempty"`
}

type Reports struct {
	Comment string   `xml:",comment"`
	Report  []string `xml:"report,omitempty"`
}

type Reporting struct {
	Comment         string         `xml:",comment"`
	ExcludeDefaults string         `xml:"excludeDefaults,omitempty"`
	OutputDirectory string         `xml:"outputDirectory,omitempty"`
	Plugins         *ReportPlugins `xml:"plugins,omitempty"`
}

type ReportPlugins struct {
	Comment string         `xml:",comment"`
	Plugins []ReportPlugin `xml:"plugins,omitempty"`
}

type ReportPlugin struct {
	Comment    string `xml:",comment"`
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	// Multiple specifications of a set of reports, each having (possibly) different configuration.
	// This is the reporting parallel to an <code>execution</code> in the build.
	ReportSets    *ReportSets `xml:"reportSets,omitempty"`
	Inherited     string      `xml:"inherited,omitempty"`
	Configuration *DOM        `xml:"configuration,omitempty"`
}

type ReportSets struct {
	Comment   string      `xml:",comment"`
	ReportSet []ReportSet `xml:"reportSet,omitempty"`
}

type ReportSet struct {
	Comment       string   `xml:",comment"`
	Id            string   `xml:"id,omitempty"`
	Reports       *Reports `xml:"reports,omitempty"`
	Inherited     string   `xml:"inherited,omitempty"`
	Configuration *DOM     `xml:"configuration,omitempty"`
}

type PluginManagement struct {
	Comment string   `xml:",comment"`
	Plugins *Plugins `xml:"plugins,omitempty"`
}

type Plugins struct {
	Comment string   `xml:",comment"`
	Plugin  []Plugin `xml:"plugin,omitempty"`
}

type Plugin struct {
	Comment      string        `xml:",comment"`
	GroupId      string        `xml:"groupId,omitempty"`
	ArtifactId   string        `xml:"artifactId,omitempty"`
	Version      string        `xml:"version,omitempty"`
	Extensions   bool          `xml:"extensions,omitempty"`
	Executions   *Executions   `xml:"executions,omitempty"`
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
	// Deprecated: Not used by Maven. Use Goals within Execution instead
	Goals         *Goals `xml:"goals,omitempty"`
	Inherited     string `xml:"inherited,omitempty"`
	Configuration *DOM   `xml:"configuration,omitempty"`
}

type Executions struct {
	Comment   string      `xml:",comment"`
	Execution []Execution `xml:"execution,omitempty"`
}

type Execution struct {
	Comment string `xml:",comment"`
	Id      string `xml:"id,omitempty"`
	Phase   string `xml:"phase,omitempty"`
	Goals   *Goals `xml:"goals,omitempty"`
}

type Goals struct {
	Comment string   `xml:",comment"`
	Goal    []string `xml:"goal,omitempty"`
}

type Resources struct {
	Comment string `xml:",comment"`
	/*
		Describe the resource target path. The path is relative to the target/classes directory

		IE: ${project.build.outputDirectory}
	*/
	TargetPath string    `xml:"targetPath,omitempty"`
	Filtering  string    `xml:"filtering,omitempty"`
	Directory  string    `xml:"directory,omitempty"`
	Includes   *Includes `xml:"includes,omitempty"`
	Excludes   *Excludes `xml:"excludes,omitempty"`
}

type Includes struct {
	Comment string   `xml:",comment"`
	Include []string `xml:"include,omitempty"`
}

type Excludes struct {
	Comment string   `xml:",comment"`
	Exclude []string `xml:"exclude,omitempty"`
}

type Filters struct {
	Comment string `xml:",comment"`
}

type BuildBase struct {
	Comment          string            `xml:",comment"`
	DefaultGoal      string            `xml:"defaultGoal,omitempty"`
	Resources        *Resources        `xml:"resources,omitempty"`
	TestResources    *Resources        `xml:"testResources,omitempty"`
	Directory        string            `xml:"directory,omitempty"`
	FinalName        string            `xml:"finalName,omitempty"`
	Filters          *Filters          `xml:"filters,omitempty"`
	PluginManagement *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugins          *Plugins          `xml:"plugins,omitempty"`
}

type Build struct {
	Comment               string      `xml:",comment"`
	SourceDirectory       string      `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string      `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string      `xml:"testSourceDirectory,omitempty"`
	OutputDirectory       string      `xml:"outputDirectory,omitempty"`
	TestOutputDirectory   string      `xml:"testOutputDirectory,omitempty"`
	Extensions            *Extensions `xml:"extensions,omitempty"`
	BuildBase
}

type Extensions struct {
	Comment   string      `xml:",comment"`
	Extension []Extension `xml:"extension,omitempty"`
}

type Extension struct {
	Comment    string `xml:",comment"`
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type Profiles struct {
	Comment string    `xml:",comment"`
	Profile []Profile `xml:"profile,omitempty"`
}

// Modifications to the build process which is activated based on environmental parameters or command line arguments.
type Profile struct {
	Comment                string                  `xml:",comment"`
	Id                     string                  `xml:"id,omitempty"`
	Activation             *Activation             `xml:"activation,omitempty"`
	Build                  *BuildBase              `xml:"build,omitempty"`
	Modules                *Modules                `xml:"modules,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *DOM                    `xml:"properties,omitempty"`
	Dependencies           *Dependencies           `xml:"dependencies,omitempty"`
	Repositories           *Repositories           `xml:"repositories,omitempty"`
	PluginRepositories     *PluginRepositories     `xml:"pluginRepositories,omitempty"`
	// Deprecated: Not used by Maven
	Reports   *Reports   `xml:"reports,omitempty"`
	Reporting *Reporting `xml:"reporting,omitempty"`
}

type Activation struct {
	Comment         string              `xml:",comment"`
	ActiveByDefault bool                `xml:"activeByDefault,omitempty"`
	JDK             string              `xml:"jdk,omitempty"`
	OS              *ActivationOS       `xml:"os,omitempty"`
	Property        *ActivationProperty `xml:"property,omitempty"`
	File            *ActivationFile     `xml:"file,omitempty"`
}

type ActivationProperty struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Value   string `xml:"value,omitempty"`
}

type ActivationOS struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Family  string `xml:"family,omitempty"`
	Arch    string `xml:"arch,omitempty"`
	Version string `xml:"version,omitempty"`
}

type ActivationFile struct {
	Comment string `xml:",comment"`
	// The name of the file that must be missing to activate the profile.
	Missing string `xml:"missing,omitempty"`
	// The name of the file that must exist to activate the profile.
	Exists string `xml:"exists,omitempty"`
}

type Properties struct {
	Comment string            `xml:",comment"`
	Fields  map[string]string `xml:"-"`
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Fields = make(map[string]string)

	type element struct {
		XMLName xml.Name
		Value   string `xml:",chardata"`
	}

	for {
		tok, err := d.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			var el element
			if err := d.DecodeElement(&el, &t); err != nil {
				return err
			}
			key := el.XMLName.Local
			value := strings.TrimSpace(el.Value)
			p.Fields[key] = value
		}
	}

	return nil
}

func (p *Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Start the root element
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Iterate through the map and create XML elements for each key-value pair
	for key, value := range p.Fields {
		element := xml.StartElement{Name: xml.Name{Local: key}}
		if err := e.EncodeElement(value, element); err != nil {
			return err
		}
	}

	// Close the root element
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

type DOM struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:"-"`
	Value    string     `xml:",innerxml"`
	Children []DOM      `xml:",any"`
}

func (a *DOM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	var children []DOM

	a.XMLName.Local = start.Name.Local
	// a.XMLName.Space = start.Name.Space

	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch el := t.(type) {
		case xml.StartElement:
			{
				var child DOM
				if err := d.DecodeElement(&child, &el); err != nil {
					return err
				}
				children = append(children, child)
			}
		case xml.CharData:
			{
				value = string(el)
			}
		}
	}

	a.Value = strings.TrimSpace(value)
	a.Children = children

	return nil
}

func (a *DOM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = a.XMLName
	start.Attr = nil // NOTE: just omitting the attrs

	// Encode the start of the element
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, child := range a.Children {
		if err := e.Encode(child); err != nil {
			return err
		}
	}

	if a.Value != "" {
		if err := e.EncodeToken(xml.CharData([]byte(a.Value))); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: a.XMLName})
}
