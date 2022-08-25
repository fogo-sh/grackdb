import { FaExternalLinkAlt, FaGlobe } from "react-icons/fa";

export function SiteReference({ site, hasLink = false }) {
  return (
    <div className="flex my-2 items-center">
      <FaGlobe className="mr-1" />
      {hasLink ? (
        <>
          <a href={site.url} className="flex items-center">
            {site.url}
            <FaExternalLinkAlt className="ml-1 text-xs" />
          </a>
        </>
      ) : (
        site.url
      )}
    </div>
  );
}
